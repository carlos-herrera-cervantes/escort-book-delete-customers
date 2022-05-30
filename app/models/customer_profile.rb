# frozen_string_literal: true

# profile table for customer_profile DB
class CustomerProfile < ActiveRecord::Base
  self.abstract_class = true
  self.table_name = 'profile'

  connects_to database: { writing: :customer_db }
end
