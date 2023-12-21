#![allow(unused)]

use actix::{Actor, StreamHandler};
use actix_web::{HttpServer, App, web, get, HttpResponse, Responder, HttpRequest, post};
use actix_web_actors::ws;

#[get("/")]
async fn index() -> actix_web::Result<actix_files::NamedFile> {
    Ok(actix_files::NamedFile::open("../static/index.html")?)
}

#[get("/ws")]
async fn ws_handle() -> actix_web::Result<actix_files::NamedFile> {
    Ok(actix_files::NamedFile::open("../static/ws.html")?)
}

#[get("/echo")]
async fn echo(req: HttpRequest, body: String) -> String {
    format!("{} {}\nbody: {}", req.method(), req.uri(), body )
}

#[post("/echo")]
async fn post_echo(req: HttpRequest, body: String) -> String {
    format!("{} {}\nbody: {}", req.method(), req.uri(), body )
}

struct MyWs;

impl Actor for MyWs {
    type Context = ws::WebsocketContext<Self>;
}

/// Handler for ws::Message message
impl StreamHandler<Result<ws::Message, ws::ProtocolError>> for MyWs {
    fn handle(&mut self, msg: Result<ws::Message, ws::ProtocolError>, ctx: &mut Self::Context) {
        match msg {
            Ok(ws::Message::Text(text)) => ctx.text(text),
            _ => (),
        }
    }
}

async fn wsapi(req: HttpRequest, stream: web::Payload) -> Result<HttpResponse, actix_web::Error> {
    ws::start(MyWs {}, &req, stream)
}

#[actix_web::main]
async fn main() {
    let _ = HttpServer::new(|| {
        App::new()
            .service(index)
            .service(ws_handle)
            .service(echo)
            .service(post_echo)
            .service(actix_files::Files::new("/", "../static"))
            .route("/api/ws", web::get().to(wsapi))
    })
    .bind("127.0.0.1:8000")
    .unwrap()
    .run()
    .await;
}





