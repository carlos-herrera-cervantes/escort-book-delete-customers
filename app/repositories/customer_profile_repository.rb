# frozen_string_literal: true

# This repository apply operations on profile table for customer_profile DB
class CustomerProfileRepository
  def find
    CustomerProfile.all
  end

  def destroy_all_by_filter(filter)
    CustomerProfile.where(filter).destroy_all
  end
end
