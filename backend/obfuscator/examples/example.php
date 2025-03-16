<?php

CONST CONSTANT_NAME = 'CONSTANT_NAME';
const constantName = "constantName";
define("constant_name", "constant_name");
Const ConstantName = "ConstantName";

$variableName1 = "VARIABLE_NAME_1";
$variableName2 = "VARIABLE_NAME_2";
$variableName3 = null;
$variableName4 = null;
$variableName5 = null;
$variableName6 = null;
$variableName7 = null;
$variableName8 = null;
$variableName9 = null;
$variableName10 = null;

// REGULAR FUNCTION WITHOUT RETURN
function functionName1()
{
    # constantName
    #### echo "DECEPTION: constantNameCONSTANT_NAME constant_name-ConstantName\n";
    echo "DECEPTION: constantNameCONSTANT_NAME constant_name-ConstantName\n"; // Just string
    echo "DECEPTION: " . constantName . " " . CONSTANT_NAME . "\n"; // This is constant, bro
    echo 'DECEPTION: ${CONSTANT_NAME}' . "\n"; // String
    echo "DECEPTION: ${constant_name}" . "\n"; // String again
    echo constantName . " " . CONSTANT_NAME . " " . constant_name . " " . ConstantName . "\n";
}

// ARROW FUNCTION WITHOUT RETURN (PHP 7.4+ menggunakan fungsi anonim)
$functionName2 = function () use ($variableName6, $variableName5, $variableName2, $variableName3, $variableName4) {
    // variableName6
    // echo "DECEPTION: variableName1 variableName6variableName4\n";
    echo "DECEPTION: variableName3 variableName2variableName5\n"; // Just string
    echo "DECEPTION: " . $variableName6 . " " . $variableName2 . "\n"; // This is variable, bro
    echo 'DECEPTION: ${variableName6}' . "\n"; // String
    echo "DECEPTION: ${variableName6}" . "\n"; // String again
    echo $variableName6 . " " . $variableName5 . " " . $variableName2 . " " . $variableName3 . " " . $variableName4 . "\n";
};

/**
 * functionName3 - regular function to sum parameters
 * @param int $params1
 * @param int $params2
 * @return int
 */
function functionName3(int $params1, int $params2): int
{
    return $params1 + $params2;
}

/**
 * functionName4 - anonymous function to sum parameters
 * @param int $params1
 * @param int $params2
 * @return int
 */
$functionName4 = fn($params1, $params2) => $params1 + $params2;

// ARROW FUNCTION WITHOUT PARENTHESES AND BRACKETS (PHP 7.4+)
$functionName5 = fn($params) => print($params . "\n");

// REGULAR FUNCTION WITH SPREAD PARAMETER
function functionName6(...$params)
{
    foreach ($params as $param) {
        echo $param . "\n";
    }
}

functionName1();
$functionName2();
$functionName3Result = functionName3(10, 20);
$functionName4Result = $functionName4(20, 10);
$functionName5("DONE");
functionName6("A", "N", "J", "A", "Y");

/** LAST STATEMENT **/
if ($functionName3Result === $functionName4Result) {
    echo "FUNCTION RESULT IS IDENTICAL\n";
}

?>
