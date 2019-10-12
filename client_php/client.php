<?php
require_once("vendor/autoload.php");
require_once("../proto_out/Helloworld/HelloRequest.php");
require_once("../proto_out/Helloworld/HelloResponse.php");
require_once("../proto_out/Helloworld/GreeterClient.php");
require_once("../proto_out/GPBMetadata/Proto/Helloworld.php");
try{
	$client = new Helloworld\GreeterClient("127.0.0.1:50051",[
			'credentials' => Grpc\ChannelCredentials::createInsecure()
	]);
	$request = new Helloworld\HelloRequest();
	$name="hetal";
	$request->setName($name);
	//print_r($request);
	list($reply, $status) = $client->SayHello($request)->wait();
	$message = $reply->getMessage();
	//print_r($reply);
	print_r($message);
	//echo strlen($message);
	//echo "\n";
echo "\n\n\nxxx\n";
	$data = get("http://127.0.0.1:50051/v1/helloworld/xxx");
echo $data;
echo strlen($data);
	$response = new Helloworld\HelloResponse($data);
	$message = $response->getMessage();
	//print_r($reply);
	print_r($message);
echo "\n\n\nxxx\n";

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
