

function tallOrWide(frac){
    if (! frac) {
        frac = 1.3
    }
    
    var torwide = function(){
        var mDiv = document.getElementById("main-area");
        if (window.innerWidth > window.innerHeight * frac){
            console.log("Wide");
            mDiv.className = "main-wide";
        }else {
            console.log("Tall");
            mDiv.className = "main-tall";
        }
    }
    
    window.addEventListener('resize',torwide);

    torwide()
    return function(n){
        if (n === undefined) {
            return n
        }
        frac = n
        torwide()
    }
}


TallFrac = tallOrWide();
 
