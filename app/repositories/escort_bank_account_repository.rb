# frozen_string_literal: true

# This repository apply operations on escort_bank_account collection for escort-book-payment DB
class EscortBankAccountRepository
  def destroy_all_by_filter(filter)
    EscortBankAccount.destroy_all filter
  end
end
