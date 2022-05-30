# frozen_string_literal: true

class CustomerBankAccount
  include Mongoid::Document
  store_in database: 'escort-book-payment', collection: 'customerbankaccounts'
end
