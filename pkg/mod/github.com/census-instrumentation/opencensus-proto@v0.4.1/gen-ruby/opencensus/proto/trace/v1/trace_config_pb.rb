# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: opencensus/proto/trace/v1/trace_config.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("opencensus/proto/trace/v1/trace_config.proto", :syntax => :proto3) do
    add_message "opencensus.proto.trace.v1.TraceConfig" do
      optional :max_number_of_attributes, :int64, 4
      optional :max_number_of_annotations, :int64, 5
      optional :max_number_of_message_events, :int64, 6
      optional :max_number_of_links, :int64, 7
      oneof :sampler do
        optional :probability_sampler, :message, 1, "opencensus.proto.trace.v1.ProbabilitySampler"
        optional :constant_sampler, :message, 2, "opencensus.proto.trace.v1.ConstantSampler"
        optional :rate_limiting_sampler, :message, 3, "opencensus.proto.trace.v1.RateLimitingSampler"
      end
    end
    add_message "opencensus.proto.trace.v1.ProbabilitySampler" do
      optional :samplingProbability, :double, 1
    end
    add_message "opencensus.proto.trace.v1.ConstantSampler" do
      optional :decision, :enum, 1, "opencensus.proto.trace.v1.ConstantSampler.ConstantDecision"
    end
    add_enum "opencensus.proto.trace.v1.ConstantSampler.ConstantDecision" do
      value :ALWAYS_OFF, 0
      value :ALWAYS_ON, 1
      value :ALWAYS_PARENT, 2
    end
    add_message "opencensus.proto.trace.v1.RateLimitingSampler" do
      optional :qps, :int64, 1
    end
  end
end

module OpenCensus
  module Proto
    module Trace
      module V1
        TraceConfig = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("opencensus.proto.trace.v1.TraceConfig").msgclass
        ProbabilitySampler = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("opencensus.proto.trace.v1.ProbabilitySampler").msgclass
        ConstantSampler = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("opencensus.proto.trace.v1.ConstantSampler").msgclass
        ConstantSampler::ConstantDecision = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("opencensus.proto.trace.v1.ConstantSampler.ConstantDecision").enummodule
        RateLimitingSampler = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("opencensus.proto.trace.v1.RateLimitingSampler").msgclass
      end
    end
  end
end