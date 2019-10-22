<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Order;

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
     * @param \Order\AddRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Add(\Order\AddRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/order.Info/Add',
        $argument,
        ['\Order\AddResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Order\ListRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function List(\Order\ListRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/order.Info/List',
        $argument,
        ['\Order\ListResponse', 'decode'],
        $metadata, $options);
    }

}
