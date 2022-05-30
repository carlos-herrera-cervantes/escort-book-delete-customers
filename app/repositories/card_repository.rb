# frozen_string_literal: true

# This repository apply operations on cards collection for escort-book-payment DB
class CardRepository
  def destroy_all_by_filter(filter)
    Card.destroy_all filter
  end
end
