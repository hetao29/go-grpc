<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Order\Info;

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
     * @param \Order\Info\AddRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Order\Info\AddResponse
     */
    public function Add(\Order\Info\AddRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/Order.Info.Info/Add',
        $argument,
        ['\Order\Info\AddResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Order\Info\ListRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Order\Info\ListResponse
     */
    public function List(\Order\Info\ListRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/Order.Info.Info/List',
        $argument,
        ['\Order\Info\ListResponse', 'decode'],
        $metadata, $options);
    }

}
