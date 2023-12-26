
const proto = window.location.protocol == "https:" ? "wss" : "ws"

const ws = new WebSocket(`${proto}://${window.location.host}/api/ws`);

ws.onopen = onopen;
ws.onmessage = wsmsg;
ws.onclose = wscls;

let id = "0"

document.querySelector("#send").addEventListener("click", () => {
  if (ws.readyState == ws.OPEN) {
    ws.send("Oof from " + id)
  } else {
    console.log("Status:", ws.readyState)
  }
})

/**
  * @arg {Event} event
  */
function onopen(event){
  console.log("WS Open", event.type)
}

/**
  * @arg {MessageEvent<any>} event
  */
function wsmsg(event){
  if (event.data.startsWith("id:")) {
    id = event.data.slice(4)
    document.querySelector("h1").innerHTML = `WebSocket ${id}`
    return 
  }

  console.log("WS Message", event.type, event.data)

  let msg = document.querySelector("#msg")
  msg.innerHTML = msg.innerHTML + `<div>${event.data}</div>`
}


/**
  * @arg {CloseEvent} event
  */
function wscls(event){
  console.log("WS Close", event.type, event.reason, event.code)
}
