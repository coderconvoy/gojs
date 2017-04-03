function tabboxtabber(box,item,buttons){
    return function(){
        console.log("Called");
        box.innerHTML = item;
        for( i in buttons) {
            buttons[i].className = "";
        }
        this.className="tabselect";
    }
}

//takes an array of objects{tabname,content}
function tabbox(tabsarr,divID){
    if (divID === undefined) {
        divID = "tabbox";
    }

    buttons = [];

    var mainDiv = document.createElement("div");
    mainDiv.id = divID
    var cDiv = document.createElement("div");
    mainDiv.appendChild(cDiv)
    var buttsDiv = document.createElement("div");
    mainDiv.appendChild(buttsDiv)
    for (var i = 0; i < tabsarr.length; i++){
        
        butt = document.createElement("button");
        butt.onclick=tabboxtabber(cDiv,tabsarr[i].content,buttons);
        butt.innerHTML = tabsarr[i].tabname;

        buttsDiv.appendChild(butt);
        buttons.push(butt);
         
    }

    cDiv.innerHTML = tabsarr[0].content;
    buttons[0].className = "tabselect";
    return mainDiv


}





