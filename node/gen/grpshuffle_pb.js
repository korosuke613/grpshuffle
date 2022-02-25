// source: grpshuffle.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.grpshuffle.Combination', null, global);
goog.exportSymbol('proto.grpshuffle.ShuffleRequest', null, global);
goog.exportSymbol('proto.grpshuffle.ShuffleResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.grpshuffle.Combination = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.grpshuffle.Combination.repeatedFields_, null);
};
goog.inherits(proto.grpshuffle.Combination, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.grpshuffle.Combination.displayName = 'proto.grpshuffle.Combination';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.grpshuffle.ShuffleRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.grpshuffle.ShuffleRequest.repeatedFields_, null);
};
goog.inherits(proto.grpshuffle.ShuffleRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.grpshuffle.ShuffleRequest.displayName = 'proto.grpshuffle.ShuffleRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.grpshuffle.ShuffleResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.grpshuffle.ShuffleResponse.repeatedFields_, null);
};
goog.inherits(proto.grpshuffle.ShuffleResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.grpshuffle.ShuffleResponse.displayName = 'proto.grpshuffle.ShuffleResponse';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.grpshuffle.Combination.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.grpshuffle.Combination.prototype.toObject = function(opt_includeInstance) {
  return proto.grpshuffle.Combination.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.grpshuffle.Combination} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpshuffle.Combination.toObject = function(includeInstance, msg) {
  var f, obj = {
    targetsList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.grpshuffle.Combination}
 */
proto.grpshuffle.Combination.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.grpshuffle.Combination;
  return proto.grpshuffle.Combination.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.grpshuffle.Combination} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.grpshuffle.Combination}
 */
proto.grpshuffle.Combination.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addTargets(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.grpshuffle.Combination.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.grpshuffle.Combination.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.grpshuffle.Combination} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpshuffle.Combination.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTargetsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
};


/**
 * repeated string targets = 1;
 * @return {!Array<string>}
 */
proto.grpshuffle.Combination.prototype.getTargetsList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.grpshuffle.Combination} returns this
 */
proto.grpshuffle.Combination.prototype.setTargetsList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.grpshuffle.Combination} returns this
 */
proto.grpshuffle.Combination.prototype.addTargets = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.grpshuffle.Combination} returns this
 */
proto.grpshuffle.Combination.prototype.clearTargetsList = function() {
  return this.setTargetsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.grpshuffle.ShuffleRequest.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.grpshuffle.ShuffleRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.grpshuffle.ShuffleRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.grpshuffle.ShuffleRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpshuffle.ShuffleRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    targetsList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    divide: jspb.Message.getFieldWithDefault(msg, 4, 0),
    sequential: jspb.Message.getBooleanFieldWithDefault(msg, 3, false)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.grpshuffle.ShuffleRequest}
 */
proto.grpshuffle.ShuffleRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.grpshuffle.ShuffleRequest;
  return proto.grpshuffle.ShuffleRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.grpshuffle.ShuffleRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.grpshuffle.ShuffleRequest}
 */
proto.grpshuffle.ShuffleRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addTargets(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setDivide(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setSequential(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.grpshuffle.ShuffleRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.grpshuffle.ShuffleRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.grpshuffle.ShuffleRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpshuffle.ShuffleRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTargetsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getDivide();
  if (f !== 0) {
    writer.writeUint64(
      4,
      f
    );
  }
  f = message.getSequential();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
};


/**
 * repeated string targets = 1;
 * @return {!Array<string>}
 */
proto.grpshuffle.ShuffleRequest.prototype.getTargetsList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.grpshuffle.ShuffleRequest} returns this
 */
proto.grpshuffle.ShuffleRequest.prototype.setTargetsList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.grpshuffle.ShuffleRequest} returns this
 */
proto.grpshuffle.ShuffleRequest.prototype.addTargets = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.grpshuffle.ShuffleRequest} returns this
 */
proto.grpshuffle.ShuffleRequest.prototype.clearTargetsList = function() {
  return this.setTargetsList([]);
};


/**
 * optional uint64 divide = 4;
 * @return {number}
 */
proto.grpshuffle.ShuffleRequest.prototype.getDivide = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.grpshuffle.ShuffleRequest} returns this
 */
proto.grpshuffle.ShuffleRequest.prototype.setDivide = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional bool sequential = 3;
 * @return {boolean}
 */
proto.grpshuffle.ShuffleRequest.prototype.getSequential = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 3, false));
};


/**
 * @param {boolean} value
 * @return {!proto.grpshuffle.ShuffleRequest} returns this
 */
proto.grpshuffle.ShuffleRequest.prototype.setSequential = function(value) {
  return jspb.Message.setProto3BooleanField(this, 3, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.grpshuffle.ShuffleResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.grpshuffle.ShuffleResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.grpshuffle.ShuffleResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.grpshuffle.ShuffleResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpshuffle.ShuffleResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    combinationsList: jspb.Message.toObjectList(msg.getCombinationsList(),
    proto.grpshuffle.Combination.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.grpshuffle.ShuffleResponse}
 */
proto.grpshuffle.ShuffleResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.grpshuffle.ShuffleResponse;
  return proto.grpshuffle.ShuffleResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.grpshuffle.ShuffleResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.grpshuffle.ShuffleResponse}
 */
proto.grpshuffle.ShuffleResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.grpshuffle.Combination;
      reader.readMessage(value,proto.grpshuffle.Combination.deserializeBinaryFromReader);
      msg.addCombinations(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.grpshuffle.ShuffleResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.grpshuffle.ShuffleResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.grpshuffle.ShuffleResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpshuffle.ShuffleResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCombinationsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.grpshuffle.Combination.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Combination combinations = 1;
 * @return {!Array<!proto.grpshuffle.Combination>}
 */
proto.grpshuffle.ShuffleResponse.prototype.getCombinationsList = function() {
  return /** @type{!Array<!proto.grpshuffle.Combination>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.grpshuffle.Combination, 1));
};


/**
 * @param {!Array<!proto.grpshuffle.Combination>} value
 * @return {!proto.grpshuffle.ShuffleResponse} returns this
*/
proto.grpshuffle.ShuffleResponse.prototype.setCombinationsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.grpshuffle.Combination=} opt_value
 * @param {number=} opt_index
 * @return {!proto.grpshuffle.Combination}
 */
proto.grpshuffle.ShuffleResponse.prototype.addCombinations = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.grpshuffle.Combination, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.grpshuffle.ShuffleResponse} returns this
 */
proto.grpshuffle.ShuffleResponse.prototype.clearCombinationsList = function() {
  return this.setCombinationsList([]);
};


goog.object.extend(exports, proto.grpshuffle);
