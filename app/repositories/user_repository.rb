# frozen_string_literal: true

# This repository apply operations on user collection for escort-book-authorizer DB
class UserRepository
  def destroy_all_by_filter(filter)
    User.destroy_all filter
  end
end
