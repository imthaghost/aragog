# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: aragog/aragog.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x13\x61ragog/aragog.proto\x12\x1cgithub.com.scaletrade.aragog\"\x0b\n\tHealthReq\"\x1c\n\nHealthResp\x12\x0e\n\x06status\x18\x01 \x01(\x05\"\x1b\n\x07UserReq\x12\x10\n\x08username\x18\x01 \x01(\t\"\x1c\n\nInviteResp\x12\x0e\n\x06status\x18\x01 \x01(\x05\"\x1c\n\nRemoveResp\x12\x0e\n\x06status\x18\x01 \x01(\x05\x32\xa8\x02\n\x06\x41ragog\x12`\n\x0bHealthCheck\x12\'.github.com.scaletrade.aragog.HealthReq\x1a(.github.com.scaletrade.aragog.HealthResp\x12]\n\nInviteUser\x12%.github.com.scaletrade.aragog.UserReq\x1a(.github.com.scaletrade.aragog.InviteResp\x12]\n\nRemoveUser\x12%.github.com.scaletrade.aragog.UserReq\x1a(.github.com.scaletrade.aragog.RemoveRespB\x0bZ\t./;aragogb\x06proto3')



_HEALTHREQ = DESCRIPTOR.message_types_by_name['HealthReq']
_HEALTHRESP = DESCRIPTOR.message_types_by_name['HealthResp']
_USERREQ = DESCRIPTOR.message_types_by_name['UserReq']
_INVITERESP = DESCRIPTOR.message_types_by_name['InviteResp']
_REMOVERESP = DESCRIPTOR.message_types_by_name['RemoveResp']
HealthReq = _reflection.GeneratedProtocolMessageType('HealthReq', (_message.Message,), {
  'DESCRIPTOR' : _HEALTHREQ,
  '__module__' : 'aragog.aragog_pb2'
  # @@protoc_insertion_point(class_scope:github.com.scaletrade.aragog.HealthReq)
  })
_sym_db.RegisterMessage(HealthReq)

HealthResp = _reflection.GeneratedProtocolMessageType('HealthResp', (_message.Message,), {
  'DESCRIPTOR' : _HEALTHRESP,
  '__module__' : 'aragog.aragog_pb2'
  # @@protoc_insertion_point(class_scope:github.com.scaletrade.aragog.HealthResp)
  })
_sym_db.RegisterMessage(HealthResp)

UserReq = _reflection.GeneratedProtocolMessageType('UserReq', (_message.Message,), {
  'DESCRIPTOR' : _USERREQ,
  '__module__' : 'aragog.aragog_pb2'
  # @@protoc_insertion_point(class_scope:github.com.scaletrade.aragog.UserReq)
  })
_sym_db.RegisterMessage(UserReq)

InviteResp = _reflection.GeneratedProtocolMessageType('InviteResp', (_message.Message,), {
  'DESCRIPTOR' : _INVITERESP,
  '__module__' : 'aragog.aragog_pb2'
  # @@protoc_insertion_point(class_scope:github.com.scaletrade.aragog.InviteResp)
  })
_sym_db.RegisterMessage(InviteResp)

RemoveResp = _reflection.GeneratedProtocolMessageType('RemoveResp', (_message.Message,), {
  'DESCRIPTOR' : _REMOVERESP,
  '__module__' : 'aragog.aragog_pb2'
  # @@protoc_insertion_point(class_scope:github.com.scaletrade.aragog.RemoveResp)
  })
_sym_db.RegisterMessage(RemoveResp)

_ARAGOG = DESCRIPTOR.services_by_name['Aragog']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\t./;aragog'
  _HEALTHREQ._serialized_start=53
  _HEALTHREQ._serialized_end=64
  _HEALTHRESP._serialized_start=66
  _HEALTHRESP._serialized_end=94
  _USERREQ._serialized_start=96
  _USERREQ._serialized_end=123
  _INVITERESP._serialized_start=125
  _INVITERESP._serialized_end=153
  _REMOVERESP._serialized_start=155
  _REMOVERESP._serialized_end=183
  _ARAGOG._serialized_start=186
  _ARAGOG._serialized_end=482
# @@protoc_insertion_point(module_scope)
