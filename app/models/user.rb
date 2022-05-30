# frozen_string_literal: true

class User
  include Mongoid::Document
  store_in collection: 'users'
end
