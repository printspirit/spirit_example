<?php
require_once("spirit.php");
$url=get_edit_url($_GET['subclass'], $_GET['tpid']);

if (!empty($_GET['target']) && $_GET['target']=='new') {
    header("Location: " . $url);
    return;
}
?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>A Test Client for Spirit</title>
    <link rel="stylesheet" href="style.css"/>
  </head>
  <body>
  <div>
  嵌入打印精灵到IFRAME中
  </div>
  <div class="iframe">
     <iframe src="<?php echo $url?>" />
  </div>
  </body>
</html> 


