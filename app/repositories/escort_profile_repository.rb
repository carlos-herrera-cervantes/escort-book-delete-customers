# frozen_string_literal: true

# This repository apply operations on profile table for escort_profile DB
class EscortProfileRepository
  def destroy_all_by_filter(filter)
    EscortProfile.where(filter).destroy_all
  end
end
