

setInterval(function () {

    $.get("http://127.0.0.1/read/", function(data, status){

        if(data.Success){
            for (i=0; i < data.Queue.length;i++) {
                packet = data.Queue[i]
                li = "<li data-json='"+JSON.stringify(packet)+"'>"
                li += "<mark>"+packet.Mark+"</mark>"
                li += "<layer>"+packet.Type+" "+packet.Sequence+" "+packet.Time+"</layer>"
                li += "<flow>"+packet.SrcIP.Data+":"+packet.SrcPort.Data+"->"+packet.DstIP.Data+":"+packet.DstPort.Data+"</flow>"
                li += "<div class='buttons'> <button class='packet'>Debug Packet</button> <button class='layer'>Debug "+packet.Type+"</button> </div>"
                li += "</li>"
                $("#packets").prepend(li)
                setEvents()
            }
        }
    });

},2000)
var onTable = false
jQuery(document).ready(function () {
    setEvents()

    $(document.body).bind('mouseup', function(e){
        var selection;

        if (window.getSelection) {
            selection = window.getSelection();
        } else if (document.selection) {
            selection = document.selection.createRange();
        }
        if (onTable && selection.toString().length > 1){
            hex = selection.toString().match(/\S+/g)
            for(i=0; i < hex.length; i++){

            }
            setDesc(hex.join(" ")+" <b>=</b> "+parseInt(hex.join(""),16))

        }

    });
})
function setEvents() {
    $("#packets li button").click(function () {
        json = JSON.parse( $(this).closest("li").attr("data-json") )
        if( $(this).is(".layer")){
            Visualize(json.LayerData,json.Type,true)
        }else{
            Visualize(json.Payload,json.Type,false)
        }

    })
}

function Visualize(bytes,type,isLayer) {
    $("#hex table,#octet table,#string table,#line table").html("")
    bytes = bytes.substr(1,bytes.length-1).trim().split(" ")
    var line,hex,octet,str
    linec = 0
    for(i =0; i < bytes.length; i++){

        nonPayload = ""
        if (!isLayer){
            if(type == "TCP" && i < 53){
                nonPayload = "nonPayload"
            }
            if(type == "UDP"  && i < 41){
                nonPayload = "nonPayload"
            }
            if(type == "IP" &&  i < 33){
                nonPayload = "nonPayload"
            }
        }
        if (i == 0 || i%16 == 0){
            linec++
            $("#line table").append("<tr><td>"+linec+"</td></tr>")
            $("#hex table").append("<tr></tr>")
            hex = $("#hex table tr").last()
            $("#string table").append("<tr></tr>")
            str = $("#string table tr").last()
        }
        octal = parseInt(bytes[i])
        l = octal.toString(16)
        if (l.length == 1) l = "0"+l
        hex.append("<td class='"+nonPayload+"' id='h"+i+"'>"+l+"</td>")
        l = String.fromCharCode(bytes[i])
        str.append("<td id='s"+i+"'>"+l+"</td>")

    }
    
    $("#hex table td").click(function () {
        $("#hex table td,#string table td").removeClass("active")
        $(this).addClass("active")
        hex = parseInt($(this).html(),16)
        $("#string table td#"+$(this).attr("id").replace("h","s")).addClass("active")
        setDesc("B"+$(this).attr("id").replace("h","")+", 0x"+$(this).html()+" <b>=</b> "+(hex).toString(10)+" <b>=</b> "+String.fromCharCode((hex).toString(10)))

    }).hover(function () {
        onTable = true
    },function () {
        onTable = false
    })
    
}

function setDesc(desc) {
    $("#desc").prepend("<p>"+desc+"</p>")
}