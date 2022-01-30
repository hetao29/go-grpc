<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user/info/info.proto

namespace User\Info;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>User.Info.LogoutResponse</code>
 */
class LogoutResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.common.Error error = 1;</code>
     */
    private $error = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Common\Error $error
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\User\Info\Info::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.common.Error error = 1;</code>
     * @return \Common\Error
     */
    public function getError()
    {
        return $this->error;
    }

    /**
     * Generated from protobuf field <code>.common.Error error = 1;</code>
     * @param \Common\Error $var
     * @return $this
     */
    public function setError($var)
    {
        GPBUtil::checkMessage($var, \Common\Error::class);
        $this->error = $var;

        return $this;
    }

}

