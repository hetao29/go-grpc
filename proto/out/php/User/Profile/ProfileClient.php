<?php
// GENERATED CODE -- DO NOT EDIT!

namespace User\Profile;

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
     * @param \User\Profile\UpdateRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \User\Profile\UpdateResponse
     */
    public function Update(\User\Profile\UpdateRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/User.Profile.Profile/Update',
        $argument,
        ['\User\Profile\UpdateResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \User\Profile\GetRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \User\Profile\GetResponse
     */
    public function Get(\User\Profile\GetRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/User.Profile.Profile/Get',
        $argument,
        ['\User\Profile\GetResponse', 'decode'],
        $metadata, $options);
    }

}
