# frozen_string_literal: true

require_relative '../app/services/kafka_service'

require 'rufus-scheduler'

scheduler = Rufus::Scheduler.singleton

scheduler.every '30s' do
  # TODO: implement some logic
  p 'JOB EXECUTED'
end

kafka_service = KafkaService.new
kafka_service.consumer.subscribe('user-delete-account', default_offset: :earliest)

loop do
  kafka_service.consumer.each_message do |message|
    record_value = message.value
    p "MESSAGE: #{record_value}"
  end
rescue Kafka::Error => e
  p "ERROR WHEN CONSUMING MESSAGES: #{e.message}"
end
