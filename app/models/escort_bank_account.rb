# frozen_string_literal: true

class EscortBankAccount
  include Mongoid::Document
  store_in database: 'escort-book-payment', collection: 'escortbankaccounts'
end
