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
	//call by grpc
	$client = new User\Info\InfoClient("127.0.0.1:8881",[
			'credentials' => Grpc\ChannelCredentials::createInsecure()
	]);
	$request = new User\Info\LoginRequest();
	$name="a";
	$request->setName($name);
	$request->setPassword("123456");
	//print_r($request);
	$time_start = microtime(true);
	for($i=0;$i<=10;$i++){
		list($reply,$error) = $client->login($request)->wait();
	}
	$time_end = microtime(true);
	$time = $time_end - $time_start;
	echo "Did nothing in $time seconds\n";

	$time_start = microtime(true);
	for($i=0;$i<=10;$i++){
		file_get_contents("http://www.mxiqi.com/x.php");
	}
	$time_end = microtime(true);
	$time = $time_end - $time_start;
	echo "Did nothing in $time seconds\n";

	if($reply){
		echo "Ok\n";
		print_r($reply->getInfo()->getName());
		print_r($reply->getInfo()->getUid());
	}else{
		echo "Error\n";
		print_r($error);
	}
	//exit;
	echo "\n";
	print_r($json = $request->serializeToJsonString());
	//call by restful
	//$request=[
	//	"name"=>"admin",
	//	"password"=>"123456",
	//];
	//print_r(json_encode($request));
	$data = post("http://127.0.0.1:8880/v1/user/login",$json);
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
function post($url, $data){
	$curl_handle=curl_init();
	curl_setopt($curl_handle, CURLOPT_URL,$url);
	curl_setopt($curl_handle, CURLOPT_POST, 1);
	curl_setopt($curl_handle, CURLOPT_CONNECTTIMEOUT, 10);
	curl_setopt($curl_handle, CURLOPT_POSTFIELDS, $data);
	curl_setopt($curl_handle, CURLOPT_RETURNTRANSFER, true);
	//curl_setopt($curl_handle, CURLOPT_USERAGENT, 'Your application name');
	$query = curl_exec($curl_handle);
	curl_close($curl_handle);
	return $query;
}
