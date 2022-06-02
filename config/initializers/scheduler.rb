# frozen_string_literal: true

require_relative '../../app/services/remover_service'
require_relative '../../app/repositories/access_token_repository'
require_relative '../../app/repositories/card_repository'
require_relative '../../app/repositories/customer_bank_account_repository'
require_relative '../../app/repositories/customer_profile_repository'
require_relative '../../app/repositories/escort_bank_account_repository'
require_relative '../../app/repositories/escort_profile_repository'
require_relative '../../app/repositories/scheduled_removal_repository'
require_relative '../../app/repositories/user_payment_repository'
require_relative '../../app/repositories/user_repository'

require 'rufus-scheduler'

scheduler = Rufus::Scheduler.singleton
remover_service = RemoverService.new(
    ScheduledRemovalRepository.new,
    CustomerProfileRepository.new,
    EscortProfileRepository.new,
    CardRepository.new,
    AccessTokenRepository.new,
    CustomerBankAccountRepository.new,
    EscortProfileRepository.new,
    UserPaymentRepository.new,
    UserRepository.new
)

scheduler.every '2m' do
  from = Date.today.beginning_of_day
  to = Date.today.end_of_day

  remover_service.remove_customers from, to
  remover_service.remove_escorts from, to
end
