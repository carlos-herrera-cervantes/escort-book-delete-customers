# frozen_string_literal: true

# customer removal collection for escort-book-schedulers DB
class ScheduledRemoval
  include Mongoid::Document
  include Mongoid::Pagination

  store_in collection: 'customer-removal', database: 'escort-book-schedulers'

  field :user_id, type: String
  field :user_type, type: String
  field :user_email, type: String
  field :scheduled_date, type: DateTime
  field :created_at, type: DateTime
  field :executed, type: Boolean

  before_save :set_date

  protected

  def set_date
    self.created_at = DateTime.now
    self.executed = false
  end
end
