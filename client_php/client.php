<?php
define("ROOT",dirname(__FILE__)."/");
define("ROOT_PROTO_LIB",ROOT."../proto/php/");

require_once(ROOT."vendor/autoload.php");

//注册自己的类
spl_autoload_register(function($class){
	$root = ROOT_PROTO_LIB.str_replace("\\","/",$class).".php";
	require_once($root);
});

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

	$request = new Info\LogoutRequest();
	list($status,$r) = $client->logout($request)->wait();
	var_dump($status);
	print_r($r);

	$client = new Profile\ProfileClient("127.0.0.1:50000",[
			'credentials' => Grpc\ChannelCredentials::createInsecure()
	]);

	$request = new Profile\GetRequest();
	list($status,$r) = $client->get($request)->wait();
	var_dump($status);
	print_r($r);

	$request = new Profile\UpdateRequest();
	list($status,$r) = $client->update($request)->wait();
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
