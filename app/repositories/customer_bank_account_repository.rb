# frozen_string_literal: true

# This repository apply operations on customer_bank_account collection for escort-book-payment DB
class CustomerBankAccountRepository
  def destroy_all_by_filter(filter)
    CustomerBankAccount.destroy_all filter
  end
end
