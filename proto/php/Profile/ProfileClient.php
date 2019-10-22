<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Profile;

/**
 * The greeter service definition.
 */
class ProfileClient extends \Grpc\BaseStub {

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
     * @param \Profile\UpdateRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Update(\Profile\UpdateRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/profile.Profile/Update',
        $argument,
        ['\Profile\UpdateResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Profile\GetRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Get(\Profile\GetRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/profile.Profile/Get',
        $argument,
        ['\Profile\GetResponse', 'decode'],
        $metadata, $options);
    }

}
