<?php
require_once("vendor/autoload.php");
require_once("../proto/php/GPBMetadata/User/Info/Info.php");
require_once("../proto/php/Info/LoginRequest.php");
require_once("../proto/php/Info/LoginResponse.php");
require_once("../proto/php/Info/InfoClient.php");
try{
	$client = new Info\InfoClient("127.0.0.1:50000",[
			'credentials' => Grpc\ChannelCredentials::createInsecure()
	]);
	$request = new Info\LoginRequest();
	$name="hetal";
	$request->setName($name);
	//print_r($request);
	list($status,$r) = $client->login($request)->wait();
	var_dump($status);
	print_r($r);
	$data = get("http://127.0.0.1:50001/v1/user/profile/get");
	echo $data;
	echo "\n";

}catch(Exception $e){
	print_r($e);
}
function get($url){
	$curl_handle=curl_init();
	curl_setopt($curl_handle, CURLOPT_URL,$url);
	curl_setopt($curl_handle, CURLOPT_CONNECTTIMEOUT, 10);
	curl_setopt($curl_handle, CURLOPT_RETURNTRANSFER, true);
	//curl_setopt($curl_handle, CURLOPT_USERAGENT, 'Your application name');
	$query = curl_exec($curl_handle);
	curl_close($curl_handle);
	return $query;
}
