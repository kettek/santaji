function Connection() {
    this.ws = new WebSocket("ws://"+document.location.host+"/ws")

    this.ws.onopen = this.onNetOpen.bind(this)
    this.ws.onclose = this.onNetClose.bind(this)
    this.ws.onmessage = this.onNetMessage.bind(this)

    this.isFirst = true
    this.queue = []
}

Connection.prototype.onNetOpen = function(e) {
    for (let i = 0; i < this.queue.length; i++) {
        this.ws.send(this.queue[i])
    }
}
Connection.prototype.onNetClose = function(e) {
    console.log(e)
}
Connection.prototype.onNetMessage = function(e) {
		let parts = e.data.split('\n')
		for (let i = 0; i < parts.length; i++) {
    	let data = JSON.parse(parts[i])
    	if (data.d) {
    	    this.onSantaRemove(data.d)
    	} else if (data.n) {
    	    if (this.isFirst) {
    	        this.onSantaNew(data.n, true)
    	        this.isFirst = false
    	    } else {
    	        this.onSantaNew(data.n)
    	    }
    	} else if (data.c) {
    	    this.onSantaChange(data.c,data.p)
    	}
	}
}
Connection.prototype.send = function(data) {
    if (this.ws.readyState == 0) {
        this.queue.push(JSON.stringify(data))
    } else {
        this.ws.send(JSON.stringify(data))
    }
}

module.exports = Connection
