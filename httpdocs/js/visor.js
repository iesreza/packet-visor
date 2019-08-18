

setInterval(function () {

    $.get("http://127.0.0.1/read/", function(data, status){
        console.warn(data)
    });

},2000)