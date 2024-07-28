<?php
if (preg_match('/\.(?:png|jpg|jpeg|gif|css)$/', $_SERVER["REQUEST_URI"])) {
    return false;    // serve the requested resource as-is.
} 

$uri = $_SERVER['REQUEST_URI'];

$arr=explode('?', $uri);
switch ($arr[0]) {
    case '/':
    case '/list':
        require('list.php');
        break;
    case '/edit':
        require('edit.php');
        break;
    default:    
        header("HTTP/1.0 404 Not Found");
        echo '404 Page Not Found';
        break;
}


