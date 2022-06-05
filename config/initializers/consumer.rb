# frozen_string_literal: true

require_relative '../../app/services/kafka_service'
require_relative '../../app/repositories/scheduled_removal_repository'

require 'json'

delete_user_consumer = KafkaService.new
delete_user_consumer.consumer.subscribe('user-delete-account', default_offset: :earliest)

cancel_delete_consumer =  KafkaService.new
cancel_delete_consumer.consumer.subscribe('user-active-account', default_offset: :earliest)

repository = ScheduledRemovalRepository.new

Thread.new do
  loop do
    delete_user_consumer.consumer.each_message do |message|
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

      p "SCHEDULED REMOVAL OF ACCOUNT FOR USER: #{hash['user_email']}"
    rescue Kafka::Error => e
      p "ERROR SCHEDULING REMOVAL OF ACCOUNT: #{e.message}"
    end
  end
end

Thread.new do
  loop do
    cancel_delete_consumer.consumer.each_message do |message|
      record_value = message.value
      parsed_record = JSON.parse record_value

      bson_user_id = BSON::ObjectId.from_string(parsed_record['userId'])
      filter = { user_id: bson_user_id }
      repository.destroy_all_by_filter filter

      p "CANCEL DELETION OF ACCOUNT FOR USER: #{parsed_record['userId']}"
    rescue Exception => e
      p "ERROR CANCELING DELETION OF ACCOUNT: #{e.message}"
    end
  end
end
