# frozen_string_literal: true

class UserPayment
  include Mongoid::Document
  store_in database: 'escort-book-payment', collection: 'paymentusers'
end
