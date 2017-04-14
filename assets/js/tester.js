var logger ={
    logs :[],
    log:function(s){
        this.logs.push(s);
    },
    print:function(){
        for(var i in this.logs){
            console.log(this.logs[i]);
        }
    }
}



function test(obs){
    passes = 0
    fails = 0
    for(var k in obs){
        var l = Object.create(logger);
        if (obs[k](l)){
            passes ++;    
        }else {
            fails ++;
            console.log("test fail:" + k)
            l.print();
        }
    }
    console.log("Passes = " + passes, "\nFails  = " + fails);
}

exports.exp = {
    test:test
}
