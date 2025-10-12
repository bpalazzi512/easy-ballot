import { NomineeCard, type Nomination } from "./NomineeCard";
import { useState } from "react";


export function Nominations() {
    const testNominations: Nomination[] = [
        {
          id: "nom-001",
          name: "Sarah Chen",
          picture: "https://i.pravatar.cc/150?img=1",
          description: "Outstanding leadership in the Q3 product launch. Led cross-functional teams to deliver ahead of schedule.",
          numVotes: 142,
          hasVoted: true,
          podiumPosition: 1
        },
        {
          id: "nom-002",
          name: "Marcus Johnson",
          picture: "https://i.pravatar.cc/150?img=12",
          description: "Exceptional mentorship program that helped onboard 15 new team members with 98% retention rate.",
          numVotes: 128,
          hasVoted: false,
          podiumPosition: 2
        },
        {
          id: "nom-003",
          name: "Priya Patel",
          picture: "https://i.pravatar.cc/150?img=5",
          description: "Innovative cost-saving initiative that reduced operational expenses by 23% while maintaining quality.",
          numVotes: 115,
          hasVoted: true,
          podiumPosition: 3
        },
        {
          id: "nom-004",
          name: "Alex Rivera",
          picture: "https://i.pravatar.cc/150?img=8",
          description: "Created the internal automation tool that saved the team 200+ hours per month.",
          numVotes: 89,
          hasVoted: false,
          podiumPosition: null
        },
        {
          id: "nom-005",
          name: "Emma Thompson",
          picture: "https://i.pravatar.cc/150?img=9",
          description: "Spearheaded diversity and inclusion initiatives, resulting in a 40% increase in underrepresented hires.",
          numVotes: 76,
          hasVoted: true,
          podiumPosition: null
        },
        {
          id: "nom-006",
          name: "James Kim",
          picture: "https://i.pravatar.cc/150?img=13",
          description: "Resolved critical production incident at 2 AM, preventing potential $500K in lost revenue.",
          numVotes: 64,
          hasVoted: false,
          podiumPosition: null
        },
        {
          id: "nom-007",
          name: "Olivia Martinez",
          picture: "https://i.pravatar.cc/150?img=10",
          description: "Designed and implemented the new customer feedback system, improving satisfaction scores by 35%.",
          numVotes: 52,
          hasVoted: true,
          podiumPosition: null
        },
        {
          id: "nom-008",
          name: "David Okonkwo",
          picture: "https://i.pravatar.cc/150?img=14",
          description: "Negotiated partnership with key vendor, securing 3-year contract with 18% cost reduction.",
          numVotes: 41,
          hasVoted: false,
          podiumPosition: null
        },
        {
          id: "nom-009",
          name: "Sophie Dubois",
          picture: "https://i.pravatar.cc/150?img=20",
          description: "Organized company-wide sustainability initiative, achieving carbon neutrality certification.",
          numVotes: 33,
          hasVoted: true,
          podiumPosition: null
        },
        {
          id: "nom-010",
          name: "Ryan O'Sullivan",
          picture: "https://i.pravatar.cc/150?img=15",
          description: "Developed comprehensive training program that increased team productivity by 28%.",
          numVotes: 27,
          hasVoted: false,
          podiumPosition: null
        }
      ];
    const [nominations, setNominations] = useState<Nomination[]>(testNominations);
    return (
        <div>
            <h1>Nominations</h1>
            {nominations.map((nomination) => (
                <NomineeCard key={nomination.id} {...nomination} onClickUpvote={() => {}} />
            ))}
        </div>
    )
}