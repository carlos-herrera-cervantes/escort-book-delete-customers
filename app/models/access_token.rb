# frozen_string_literal: true

class AccessToken
  include Mongoid::Document
  store_in database: 'accesstokens'
end
