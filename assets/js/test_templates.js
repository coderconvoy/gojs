var tmod = require("./template.js").exp;
var tests = require("./tester.js").exp;

function test_template(logger){
    s = "hello {{pete}}, {{sam}}";
    m = {
        "pete":"dave",
        "sam":"sammy"
    }
    res = tmod.replaceAll(s,m)
    if (res != "hello dave, sammy") {
        logger.log("No match : expected :\n"+ "hello dave, sammy\ngot:\n"+res);
        return false;
    }
    return true;
}



console.log("test_template.js");
tests.test({
    "template":test_template
})



