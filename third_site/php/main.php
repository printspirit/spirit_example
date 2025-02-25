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
	case '/new':
		require('spirit.php');
		$subclass=$_GET['subclass'];
		$target=$_GET['target'];
		$id = newLabel("TEST Label", 800, 500, 203, $subclass);
		if ($id!="") {
			header("Location: " . "edit?subclass={$subclass}&tpid={$id}&target={$target}");
			return;
		}
		break;
	case '/del':
		require('spirit.php');
		$id=$_GET['tpid'];
		delLabel($id);
		header("Location: " . "list");
		break;
		
    case '/edit':
        require('edit.php');
        break;
    default:    
        header("HTTP/1.0 404 Not Found");
        echo '404 Page Not Found';
        break;
}
