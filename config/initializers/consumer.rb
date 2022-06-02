# frozen_string_literal: true

require_relative '../../app/services/kafka_service'
require_relative '../../app/repositories/scheduled_removal_repository'

require 'json'

kafka_service = KafkaService.new
kafka_service.consumer.subscribe('user-delete-account', default_offset: :earliest)
repository = ScheduledRemovalRepository.new

Thread.new do
  loop do
    kafka_service.consumer.each_message do |message|
      record_value = message.value
      parsed_record = JSON.parse record_value

      today = Date.today
      days_to_expire = 28

      hash = {}
      hash['user_id'] = parsed_record['userId']
      hash['user_type'] = parsed_record['userType']
      hash['user_email'] = parsed_record['userEmail']
      hash['scheduled_date'] = (today + days_to_expire)

      repository.create hash

      p 'MESSAGE PROCESSED'
    rescue Kafka::Error => e
      p "ERROR PROCESSING A MESSAGE: #{e.message}"
    end
  end
end
