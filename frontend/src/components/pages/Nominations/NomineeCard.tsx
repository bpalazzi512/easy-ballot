


export interface Nomination {
  id: string;
  name: string;
  picture: string;
  description: string;
  numVotes: number;
  hasVoted: boolean;
  podiumPosition: number | null;
}

interface NominationCardProps extends Nomination {
  onClickUpvote: () => void;
}

export function NomineeCard({ id, name, picture, description, numVotes, hasVoted, podiumPosition, onClickUpvote }: NominationCardProps) {
  return (
    <div className="cursor-pointer flex flex-col gap-2">
      {podiumPosition && (
        <div className="absolute top-0 left-0 w-10 h-10 bg-red-500 rounded-full">
          {podiumPosition}
        </div>
      )}

      <div className="flex items-center gap-2">
        <img src={picture} alt={name} className="w-10 h-10 rounded-full" />
        <h3 className="text-lg font-bold">{name}</h3>
      </div>
      <p className="text-sm text-gray-500">{description}</p>
      <div>
        <p className="text-sm text-gray-500">{numVotes}</p>
        <button className="text-sm text-blue-500" onClick={onClickUpvote} disabled={hasVoted}>Second</button>
      </div>
    </div>
  )
}