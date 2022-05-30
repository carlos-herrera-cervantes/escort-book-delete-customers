# frozen_string_literal: true

require 'kafka'

# Kafka client
class KafkaService
  attr_reader :kafka

  def initialize
    @kafka = Kafka.new(seed_brokers: 'localhost:9092')
  end

  def consumer
    return @consumer unless @consumer.nil?

    @consumer = @kafka.consumer(group_id: 'escort-book-delete-customers')

    at_exit do
      @consumer.stop
    end
    @consumer
  end
end
