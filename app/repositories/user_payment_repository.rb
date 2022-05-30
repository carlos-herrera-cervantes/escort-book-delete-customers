# frozen_string_literal: true

# This repository apply operations on payment_user collection for escort-book-payment DB
class UserPaymentRepository
  def destroy_all_by_filter(filter)
    UserPayment.destroy_all filter
  end
end
