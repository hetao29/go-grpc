<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/helloworld.proto

namespace GPBMetadata\Proto;

class Helloworld
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(hex2bin(
            "0aba010a1670726f746f2f68656c6c6f776f726c642e70726f746f120a68" .
            "656c6c6f776f726c64221c0a0c48656c6c6f52657175657374120c0a046e" .
            "616d6518012001280922200a0d48656c6c6f526573706f6e7365120f0a07" .
            "6d657373616765180120012809324c0a074772656574657212410a085361" .
            "7948656c6c6f12182e68656c6c6f776f726c642e48656c6c6f5265717565" .
            "73741a192e68656c6c6f776f726c642e48656c6c6f526573706f6e736522" .
            "00620670726f746f33"
        ), true);

        static::$is_initialized = true;
    }
}

