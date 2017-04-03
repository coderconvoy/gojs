//This method serves to begin a loop on a set of images located in the same array, by fading them in and out
//Returns a function which you may call to hurry it up.
function beginSlowSwitches(ob,delay,fadeTime){
    
    var count = 0; 
    if (delay === undefined) {
        delay = 2000;
    }
    if (fadeTime === undefined) fadeTime = 100;
    for (var i = 1; i < ob.children.length; i ++ )
        ob.children[i].style.opacity = 0;

    var sTime = undefined;
    var pFrom = undefined;
    var pTo = undefined;
    var cTime = 0;

    var fade = function(t){
        if (sTime === -1 ){
            sTime = t;
        }
        if (t - sTime >= fadeTime){
            pFrom.style.opacity = 0;
            pTo.style.opacity = 1;
            sTime = undefined;
            return; 
        }

        dist = (t -sTime )/fadeTime;
        pFrom.style.opacity = 1 - dist;
        pTo.style.opacity = dist;
        requestAnimationFrame(fade);
    }
    
    var cloj = function(){
        cTime = 0;
        
        if (sTime !== undefined ) return;
        sTime = -1;
        
        pFrom = ob.children[count];
        count ++;
        if (count >= ob.children.length ){
            count = 0; 
        }
        pTo = ob.children[count]; 
        
        window.requestAnimationFrame(fade);
    }

    var pCloj = function(){
        cTime ++;
        if (cTime >= 5 ){
            cloj();
        }
    }

    setInterval(pCloj,delay/5);
    return cloj;
}



