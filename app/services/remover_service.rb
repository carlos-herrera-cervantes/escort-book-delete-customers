# frozen_string_literal: true

# This service will apply the remove operations for both escort and customer information
class RemoverService
  def initialize(
      removal_repository,
      customer_profile_repository,
      escort_profile_repository,
      card_repository,
      token_repository,
      customer_bank_repository,
      escort_bank_repository,
      payment_repository,
      user_repository
  )
    @removal_repository = removal_repository
    @customer_profile_repository = customer_profile_repository
    @escort_profile_repository = escort_profile_repository
    @card_repository = card_repository
    @access_token_repository = token_repository
    @customer_bank_repository = customer_bank_repository
    @escort_bank_repository = escort_bank_repository
    @user_payment_repository = payment_repository
    @user_repository = user_repository
  end

  def remove_customers(from, to)
    removal_filter = {
        :scheduled_date.gte => from,
        :scheduled_date.lte => to,
        :user_type => 'Customer',
        :executed => false
    }
    customers = removal_repository.find_by_filter removal_filter, 0, 1000

    if customers.length == 0
      p 'NO CUSTOMERS TO DELETE'
      return
    end

    customer_ids = customers.map(&:user_id)
    bson_customer_ids = customers.map { |c| BSON::ObjectId.from_string(c.user_id) }

    # TODO: Here we need to call the payment gateway for:
    # 1 - Delete customer bank account
    # 2 - Delete customer cards
    # Think about how to delete customer cards and bank account in batches

    profile_filter = { customer_id: customer_ids }
    customer_profile_repository.destroy_all_by_filter profile_filter

    card_filter = { customerId: { '$in': bson_customer_ids } }
    card_repository.destroy_all_by_filter card_filter
    customer_bank_repository.destroy_all_by_filter card_filter

    payment_filter = { userId: { '$in': bson_customer_ids } }
    user_payment_repository.destroy_all_by_filter payment_filter

    user_filter = { _id: { '$in': bson_customer_ids } }
    user_repository.destroy_all_by_filter user_filter

    customer_emails = customers.map(&:user_email)
    token_filter = { user: { '$in': customer_emails } }
    access_token_repository.destroy_all_by_filter token_filter

    bson_removal_ids = customers.map { |c| BSON::ObjectId.from_string(c._id) }
    removal_filter = { _id: { '$in': bson_removal_ids } }
    new_values = { '$set': { executed: true } }

    removal_repository.update_many_by_filter removal_filter, new_values

    p 'SUCCESS DELETION OF CUSTOMERS'
  end

  def remove_escorts(from, to)
    removal_filter = {
        :scheduled_date.gte => from,
        :scheduled_date.lte => to,
        :user_type => 'Escort',
        :executed => false
    }
    escorts = removal_repository.find_by_filter removal_filter, 0, 1000

    if escorts.length == 0
      p 'NO ESCORTS TO DELETE'
      return
    end

    escort_ids = escorts.map(&:user_id)
    escort_emails = escorts.map(&:user_email)
    bson_escort_ids = escorts.map { |c| BSON::ObjectId.from_string(c.user_id) }

    # TODO: Here we need to call the payment gateway for:
    # 1 - Delete the escort bank account
    # Think about how to delete bank accounts in batches

    profile_filter = { escort_id: escort_ids }
    escort_profile_repository.destroy_all_by_filter profile_filter

    payment_filter = { userId: { '$in': bson_escort_ids } }
    user_payment_repository.destroy_all_by_filter payment_filter

    user_filter = { _id: { '$in' => bson_escort_ids } }
    user_repository.destroy_all_by_filter user_filter

    token_filter = { user: { '$in' => escort_emails } }
    access_token_repository.destroy_all_by_filter token_filter

    info_filter = { escortId: { '$in' => bson_escort_ids } }
    escort_bank_repository.destroy_all_by_filter info_filter

    bson_removal_ids = escorts.map { |c| BSON::ObjectId.from_string(c._id) }
    removal_filter = { _id: { '$in': bson_removal_ids } }
    new_values = { '$set': { executed: true } }

    removal_repository.update_many_by_filter removal_filter, new_values

    p 'SUCCESS DELETION OF ESCORTS'
  end

  private

  attr_reader :removal_repository, :customer_profile_repository
  attr_reader :escort_profile_repository, :card_repository
  attr_reader :access_token_repository, :customer_bank_repository
  attr_reader :escort_bank_repository, :user_payment_repository
  attr_reader :user_repository
end
