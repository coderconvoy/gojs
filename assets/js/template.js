
function replaceAll(s,map){
    res = s
    for ( var k in map) {
        res = res.replace("{{"+k+"}}" , map[k]);
    }
    return res
}



//Export for node testing
if (typeof exports !== 'undefined'){
    exports.exp = {
        replaceAll:replaceAll
    }
}
