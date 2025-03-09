const constantName = 'constantName';
const CONSTANT_NAME = 'CONSTANT_NAME';
const constant_name = 'constant_name';
const ConstantName = 'ConstantName';

let variableName1 = 'VARIABLE_NAME_1';
var variableName2 = 'VARIABLE_NAME_2';
let variableName3, variableName4;
var variableName5, variableName6;
let variableName7,variableName8 ,variableName9 , variableName10;

// REGULAR FUNCTION WITHOUT RETURN
function functionName1 () {
    // constantName
    // console.log("DECEPTION: constantNameCONSTANT_NAME constant_name-ConstantName");
    console.log("DECEPTION: constantNameCONSTANT_NAME constant_name-ConstantName"); // Just string
    console.log(`DECEPTION: ${constantName} ${CONSTANT_NAME}`); // This is constant, bro
    console.log('DECEPTION: ${CONSTANT_NAME}'); // String
    console.log("DECEPTION: ${constant_name}"); // String again
    console.log(constantName, CONSTANT_NAME, constant_name, ConstantName);
}

// ARROW FUNCTION WITHOUT RETURN
const functionName2 = () => {
    // variableName6
    // console.log("DECEPTION: variableName1 variableName6variableName4");
    console.log("DECEPTION: variableName3 variableName2variableName5"); // Just string
    console.log(`DECEPTION: ${variableName6} ${variableName2}`); // This is variable, bro
    console.log('DECEPTION: ${variableName6}'); // String
    console.log("DECEPTION: ${variableName6}"); // String again
    console.log(variableName6, variableName5, variableName2, variableName3, variableName4);
};

/**
 * functionName3 - regular function to sum parameters
 * @param params1 {int}
 * @param params2 {int}
 * @returns {int}
 */
function functionName3 (params1, params2) {
    return params1 + params2;
}

/**
 * functionName4 - arrow function to sum parameters
 * @param params1
 * @param params2
 * @returns {*}
 */
const functionName4 = (params1, params2) => params1 + params2;

// ARROW FUNCTION WITHOUT PARENTHESES AND BRACKETS
const functionName5 = params => console.log(params);

// REGULAR FUNCTION WITH SPREAD PARAMETER
function functionName6 (...params) {
    for (const param of params) {
        console.log(param);
    }
}

functionName1();
functionName2();
const functionName3Result = functionName3(10, 20);
let functionName4Result = functionName4(20, 10);
functionName5('DONE');
functionName6('A', 'N', 'J', 'A', 'Y');

/** LAST STATEMENT **/
if (functionName3Result === functionName4Result) {
    console.log('FUNCTION RESULT IS IDENTICAL');
}


