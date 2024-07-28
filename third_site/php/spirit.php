<?php 
define("SPIRIT_HOST", "https://www.printspirit.cn");
#define("SPIRIT_HOST", "http://192.168.1.100:8059");
define("UID",  "third_test");
define("PASS", "third_test");

function getAccessToken($uid, $pass) {

	$apcuAvailabe = function_exists('apcu_enabled') && apcu_enabled();
	
	if($apcuAvailabe){
		$access_token = apcu_fetch('access_token');
		$expirt_time = apcu_fetch('expirt_time');
		if ( $access_token && $expirt_time > time() ) return $access_token;
	}

	$rc=json_decode(file_get_contents(SPIRIT_HOST . "/api/get-access-token?userid=$uid&passwd=$pass"));

	if ($rc!=NULL && $rc->rc=='OK') {
		if($apcuAvailabe){
			apcu_store('access_token', $rc->token);
			apcu_store('expirt_time',  time() + $rc->expirt);
		}
		return $rc->token;
	}	
	die("无法获取TOKEN:".$rc->errmsg);
}

function getList($subclass="") {
	$token = getAccessToken(UID, PASS);
	$rc=json_decode(file_get_contents(SPIRIT_HOST . "/api/get-label-list?token=${token}&subclass=${subclass}"), true);
	if ($rc!=NULL && $rc['rc']=='OK') 	return $rc['data'];
	return [];
}

function getContent($tpid) {
	$token = getAccessToken(UID, PASS);
	$rc=json_decode(file_get_contents(SPIRIT_HOST . "/api/get-label-content?token=${token}&tpid=${tpid}"), true);
	if ($rc!=NULL && $rc['rc']=='OK') return $rc['data'];
	return "";
}

function get_edit_url($subclass, $tpid="") {
	$token = getAccessToken(UID, PASS);
	return SPIRIT_HOST . "/third-edit?subclass=${subclass}&tpid=${tpid}&token=${token}";	
}
