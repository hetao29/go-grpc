package com.yunke.app.helloworld;

import io.grpc.Channel;
import io.grpc.ManagedChannel; 
import io.grpc.ManagedChannelBuilder;
import user.info.*;

public class Client{

	private InfoGrpc.InfoBlockingStub helloWorldServiceBlockingStub;
	public void init() {
		ManagedChannel managedChannel = ManagedChannelBuilder
			.forAddress("127.0.0.1", 50000).usePlaintext().build();

		helloWorldServiceBlockingStub =
			InfoGrpc.newBlockingStub(managedChannel);
	}

	public String Get(String name, String pwd) {
		InfoOuterClass.LoginRequest request = InfoOuterClass.LoginRequest.newBuilder().setName(name)
			.setPassword(pwd).build();

		InfoOuterClass.LoginResponse response=
			helloWorldServiceBlockingStub.login(request);

		return response.getMessage();
	}

}

