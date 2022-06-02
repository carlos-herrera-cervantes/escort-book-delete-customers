# frozen_string_literal: true

class AccessToken
  include Mongoid::Document
  store_in collection: 'accesstokens'
end
