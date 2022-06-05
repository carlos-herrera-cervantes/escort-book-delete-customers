# frozen_string_literal: true

# This repository apply operations on customer removal collection for escort-book-schedulers DB
class ScheduledRemovalRepository
  def create(document)
    created = ScheduledRemoval.new(document)

    if created.save
      created
    else
      created.errors
    end
  end

  def find_by_filter(filter, offset, limit)
    ScheduledRemoval.where(filter).paginate(page: offset, limit: limit)
  end

  def update_many_by_filter(filter, new_values)
    ScheduledRemoval.where(filter).update_all(new_values)
  end

  def destroy_all_by_filter(filter)
    ScheduledRemoval.destroy_all filter
  end
end
