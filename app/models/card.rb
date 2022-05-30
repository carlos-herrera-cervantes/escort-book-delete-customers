# frozen_string_literal: true

class Card
  include Mongoid::Document
  store_in database: 'escort-book-payment', collection: 'cards'
end
