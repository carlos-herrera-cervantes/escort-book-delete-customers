# frozen_string_literal: true

# This repository apply operations on access_token collection for escort-book-authorizer DB
class AccessTokenRepository
  def destroy_all_by_filter(filter)
    AccessToken.destroy_all filter
  end
end
