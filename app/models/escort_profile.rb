# frozen_string_literal: true

# profile table for escort_profile DB
class EscortProfile < ActiveRecord::Base
  self.abstract_class = true
  self.table_name = 'profile'

  connects_to database: { writing: :escort_db }
end