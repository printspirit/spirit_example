<?php
require_once("spirit.php");
?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>A Test Client for Spirit</title>
    <link rel="stylesheet" href="style.css"/>
  </head>
  <body>
    <div class="container">
        <div class="section">
        第三方网站嵌入打印精灵
        </div>
        <div class="section-detail">
        <p>第三方网站可以通过 get-label-list API 获取本企业的全部标签列表，调用时需要提供一个附加参数subclass，为自定的分类标识，最常见的用途是用来区分该企业下的最终用户。
        如果subclass为空，返回该帐号下的全部标签。</p>
        <p>标签可以编辑，通过target可以控制编辑时嵌入到第三方网站，或跳转到打印精灵</p>
        </div>
   
      	<table class='table'>
      	<thead>
      	<tr><th>标签名</th><th>子分类</th><th>缩略图</th><th/></tr>
      	</thead>
      	<tbody>
      	<?php 
		    $lst = getList("user1");
      		foreach( $lst as $l) {
	      		?>
	      		<tr>
	      		    <td><?=$l['name']?></td>
	      		    <td><?=$l['subclass']?></td>
	      		    <td><img class="thumb" src='<?=SPIRIT_HOST . "/utils/thumb?id=${l['id']}"?>'</td>
	      		    <td>
	      		        <a href='<?="edit?subclass=${l['subclass']}&tpid=${l['id']}"?>'><button>编辑(嵌入)</button></a>
	      		        <a href='<?="edit?subclass=${l['subclass']}&tpid=${l['id']}&target=new"?>'><button>编辑(跳转)</button></a>
						<a href='<?="del?tpid=${l['id']}"?>'><button>删除</button></a>
	      		    </td>
	      		</tr>
	    <?php
	      	}
      	?>
      	</tbody>	
      	</table>
        <div class="section">
        新建标签
        </div>
        <div class="section-detail">
        新建标签时，需要制定subclass. 通过target可以控制编辑时嵌入到第三方网站，或跳转到打印精灵
        </div>      	
      	<div class="section-detail">
      		<a href='edit?subclass=user1'><button>新增标签(嵌入)</button></a>
      		<a href='edit?subclass=user1&target=new'><button>新增标签(跳转)</button></a>
			<a href='new?subclass=user1'><button>先创建后修改(嵌入)</button></a>
      		<a href='new?subclass=user1&target=new'><button>先创建后修改(跳转)</button></a>
      	</div>
  	 </div>
  </body>
</html> 
