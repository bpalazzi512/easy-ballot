import flame from '../../../assets/flame.svg';


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

  const shadowStyle = podiumPosition ? { boxShadow: '0 0 10px rgba(254, 122, 96, 0.75)' } : {boxShadow: '0 0 10px rgba(233, 233, 233, 1)'};
  return (
    <div className="flex flex-col gap-2 relative w-fit max-w-4/5 rounded-md p-3" style={shadowStyle}>
      {podiumPosition && (
        <div className="w-fit absolute -top-2/12 -right-[6%]">
          <img src={flame} alt="flame" className="w-12 h-12" />
          <p className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-4/12 text-white">
            {podiumPosition}
          </p>

        </div>
      )}

      <div className="flex items-center gap-2">
        <img src={picture} alt={name} className="w-10 h-10 rounded-full" />
        <h3 className="text-lg">{name}</h3>
      </div>
      <p className="text-sm text-gray-500">{description}</p>
      <div className="w-full flex items-center justify-end gap-2">
        <p className="text-sm text-gray-500">{numVotes}</p>
        <button className="text-sm text-white bg-theme-red rounded-md px-3 py-1 cursor-pointer" onClick={onClickUpvote} disabled={hasVoted}>Second</button>
      </div>
    </div>
  )
}