<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Info;

/**
 * The greeter service definition.
 */
class InfoClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * Sends a greeting
     * @param \Info\LoginRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function login(\Info\LoginRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/info.Info/login',
        $argument,
        ['\Info\LoginResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Info\LogoutRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function logout(\Info\LogoutRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/info.Info/logout',
        $argument,
        ['\Info\LogoutResponse', 'decode'],
        $metadata, $options);
    }

}
